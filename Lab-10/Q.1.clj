
;; Lab 10 - Q1
;; Clojure Core Immutable Data Structures

(println "===== LIST OPERATIONS =====")

(def my-list '(10 20 30 40))

(println "Original List:" my-list)
(println "First Element:" (first my-list))
(println "Rest of List:" (rest my-list))
(println "After Cons:" (cons 5 my-list))
(println "Original List Unchanged:" my-list)


(println "\n===== VECTOR OPERATIONS =====")

(def my-vector [10 20 30 40])

(println "Original Vector:" my-vector)
(println "After Conj:" (conj my-vector 50))
(println "Element at Index 2:" (get my-vector 2))
(println "Vector Count:" (count my-vector))
(println "Original Vector Unchanged:" my-vector)


(println "\n===== MAP OPERATIONS =====")

(def my-map {:name "Aditya" :age 19 :city "Gurugram"})

(println "Original Map:" my-map)
(println "After Assoc:" (assoc my-map :age 20))
(println "After Dissoc:" (dissoc my-map :city))
(println "Get Name:" (get my-map :name))
(println "Original Map Unchanged:" my-map)


(println "\n===== SET OPERATIONS =====")

(def my-set #{10 20 30 40})

(println "Original Set:" my-set)
(println "After Conj:" (conj my-set 50))
(println "After Disj:" (disj my-set 20))
(println "Contains 30:" (contains? my-set 30))
(println "Original Set Unchanged:" my-set)
