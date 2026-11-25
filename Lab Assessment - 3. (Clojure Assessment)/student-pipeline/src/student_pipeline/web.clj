(ns student-pipeline.web
  (:require
   [ring.adapter.jetty :as jetty]
   [cheshire.core :as json]
   [student-pipeline.database :as db]))

(defn json-response [status data]
  {:status status
   :headers {"Content-Type" "application/json"}
   :body (json/generate-string data
                              {:pretty true})})

(defn handler [request]
  (try
    (let [students (db/get-students)]

      (if (nil? students)
        (json-response 503
                       {:error "Database unavailable"})

        (case (:uri request)

          "/"
          (json-response 200
            {:message "Student Performance Data Pipeline"
             :endpoints ["/students"
                         "/passed"
                         "/failed"
                         "/statistics"]})

          "/students"
          (json-response 200 students)

          "/passed"
          (json-response 200
            (filter #(>= (:marks %) 40) students))

          "/failed"
          (json-response 200
            (filter #(< (:marks %) 40) students))

          "/statistics"
          (json-response 200
            {:total-students (count students)
             :passed (count
                       (filter #(>= (:marks %) 40)
                               students))
             :failed (count
                       (filter #(< (:marks %) 40)
                               students))
             :average-marks
             (if (empty? students)
               0.0
               (/ (double
                    (reduce + (map :marks students)))
                  (count students)))})

          (json-response 404
                         {:error "Page not found"}))))

    (catch Exception e
      (println "API error:" (.getMessage e))
      (json-response 500
                     {:error "Internal server error"}))))

(defn start-server []
  (println "Web server running at http://localhost:3000")
  (jetty/run-jetty handler
                   {:host "127.0.0.1"
                    :port 3000
                    :join? true}))