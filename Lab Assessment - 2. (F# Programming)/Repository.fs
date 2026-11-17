namespace LibraryManager

open System
open MongoDB.Bson
open MongoDB.Driver

type Repository(uri: string, databaseName: string) =
    let settings = MongoClientSettings.FromConnectionString(uri)
    do settings.ServerSelectionTimeout <- TimeSpan.FromSeconds(5.0)
    let client = new MongoClient(settings)
    let database = client.GetDatabase(databaseName)
    let collection = database.GetCollection<BsonDocument>("books")

    let filter id =
        Builders<BsonDocument>.Filter.Eq("_id", ObjectId.Parse(id))

    let document id title author status =
        BsonDocument(
            [
                BsonElement("_id", BsonObjectId(id))
                BsonElement("title", BsonString(title))
                BsonElement("author", BsonString(author))
                BsonElement("status", BsonString(Domain.statusText status))
            ]
        )

    member _.Initialize() =
        async {
            let! _ =
                database.RunCommandAsync<BsonDocument>(BsonDocument("ping", 1))
                |> Async.AwaitTask

            return ()
        }

    member _.Create(title, author) =
        async {
            match Domain.validate title author with
            | Error message -> return Error message
            | Ok(title, author) ->
                let book = document (ObjectId.GenerateNewId()) title author Available
                do! collection.InsertOneAsync(book) |> Async.AwaitTask
                return Ok()
        }

    member _.ReadAll() =
        async {
            let! documents =
                collection
                    .Find(Builders<BsonDocument>.Filter.Empty)
                    .Sort(Builders<BsonDocument>.Sort.Ascending("_id"))
                    .ToListAsync()
                |> Async.AwaitTask

            return
                documents
                |> Seq.map (fun doc ->
                    let status =
                        match doc.["status"].AsString with
                        | "Available" -> Available
                        | "Borrowed" -> Borrowed
                        | value -> failwithf "Invalid stored status: %s" value

                    {
                        Id = doc.["_id"].AsObjectId.ToString()
                        Title = doc.["title"].AsString
                        Author = doc.["author"].AsString
                        Status = status
                    })
                |> List.ofSeq
        }

    member _.Update(id, title, author, status) =
        async {
            match Domain.validate title author with
            | Error message -> return Error message
            | Ok(title, author) ->
                let replacement = document (ObjectId.Parse(id)) title author status

                let! result =
                    collection.ReplaceOneAsync(filter id, replacement, ReplaceOptions(IsUpsert = false))
                    |> Async.AwaitTask

                return
                    if result.MatchedCount = 1L then
                        Ok()
                    else
                        Error "Book not found."
        }

    member _.Delete(id: string) =
        async {
            let! result = collection.DeleteOneAsync(filter id) |> Async.AwaitTask

            return
                if result.DeletedCount = 1L then
                    Ok()
                else
                    Error "Book not found."
        }

    interface IDisposable with
        member _.Dispose() = client.Dispose()
