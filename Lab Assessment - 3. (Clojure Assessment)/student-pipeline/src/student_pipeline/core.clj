(ns student-pipeline.core
  (:require
   [student-pipeline.database :as db]
   [student-pipeline.web :as web])
  (:gen-class))

(def students
  [{:id 1 :name "Aditya" :marks 85}
   {:id 2 :name "Akansh" :marks 72}
   {:id 3 :name "Naman" :marks 92}
   {:id 4 :name "Kartikeya" :marks 35}
   {:id 5 :name "Hrishit" :marks 65}
   {:id 6 :name "Vignesh" :marks 48}])

(defn calculate-grade [marks]
  (cond
    (>= marks 90) "A+"
    (>= marks 80) "A"
    (>= marks 70) "B"
    (>= marks 60) "C"
    (>= marks 40) "D"
    :else "Fail"))

(defn process-students [data]
  (map (fn [student]
         (assoc student :grade
                (calculate-grade (:marks student))))
       data))

(defn passed-students [data]
  (filter #(>= (:marks %) 40) data))

(defn average-marks [data]
  (if (empty? data)
    0.0
    (/ (double (reduce + (map :marks data)))
       (count data))))

(defn -main [& args]
  (println "Student Performance Data Pipeline")
  (println "---------------------------------")

  (let [processed (vec (process-students students))]

    (println "\nProcessed Student Results:")
    (doseq [student processed]
      (println student))

    (println "\nTotal Students:" (count students))

    (println "Passed Students:"
             (count (passed-students students)))

    (println "Average Marks:"
             (format "%.2f" (average-marks students)))

    (println "\nSaving records to MongoDB...")

    (if (db/save-students processed)
      (do
        (println "\nReading records from MongoDB:")
        (doseq [student (db/get-students)]
          (println student))

        (println "\nStarting Ring HTTP Server...")
        (web/start-server))

      (println "Unable to start application."))))