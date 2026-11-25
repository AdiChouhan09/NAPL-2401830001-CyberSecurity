(ns student-pipeline.database
  (:import
   [com.mongodb.client MongoClients]
   [com.mongodb.client.model Filters ReplaceOptions]
   [org.bson Document]))

(def connection-string "mongodb://localhost:27017")
(def database-name "student_db")
(def collection-name "students")

(defn save-students [students]
  (try
    (with-open [client
                (MongoClients/create connection-string)]
      (let [database (.getDatabase client database-name)
            collection (.getCollection database collection-name)]

        (doseq [student students]
          (let [document (Document.)]
            (.append document "_id" (int (:id student)))
            (.append document "name" (:name student))
            (.append document "marks" (int (:marks student)))
            (.append document "grade" (:grade student))

            (.replaceOne collection
                         (Filters/eq "_id" (int (:id student)))
                         document
                         (.upsert (ReplaceOptions.) true))))

        (println "Student records saved to MongoDB successfully.")
        true))

    (catch Exception e
      (println "MongoDB write error:" (.getMessage e))
      false)))

(defn get-students []
  (try
    (with-open [client
                (MongoClients/create connection-string)]
      (let [database (.getDatabase client database-name)
            collection (.getCollection database collection-name)]

        (with-open [cursor (.iterator (.find collection))]
          (loop [results []]
            (if (.hasNext cursor)
              (let [document (.next cursor)]
                (recur
                 (conj results
                       {:id (.get document "_id")
                        :name (.getString document "name")
                        :marks (.get document "marks")
                        :grade (.getString document "grade")})))
              results)))))

    (catch Exception e
      (println "MongoDB read error:" (.getMessage e))
      nil)))