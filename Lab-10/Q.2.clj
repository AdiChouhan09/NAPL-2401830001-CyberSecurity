
;; Lab 10 - Q2
;; Sequence Transformations

(println "===== COLLECTION 1: NUMBERS =====")

(def numbers [1 2 3 4 5 6 7 8 9 10])

(println "Original Collection:" numbers)

;; Map - Square each number
(println "Map (Square):"
         (map #(* % %) numbers))

;; Filter - Select even numbers
(println "Filter (Even Numbers):"
         (filter even? numbers))

;; Reduce - Calculate sum
(println "Reduce (Sum):"
         (reduce + numbers))

;; Take - Select first five elements
(println "Take 5:"
         (take 5 numbers))

;; Drop - Remove first five elements
(println "Drop 5:"
         (drop 5 numbers))

;; Threading macro
(println "\nThreading Macro (Numbers):")

(def result1
  (->> numbers
       (map #(* % %))
       (filter even?)
       (take 3)))

(println "Squared, Filtered and First 3:" result1)


(println "\n===== COLLECTION 2: MARKS =====")

(def marks [12 15 18 21 24 27])

(println "Original Collection:" marks)

;; Map - Add 1 to each mark
(println "Map (Add 1):"
         (map inc marks))

;; Filter - Select marks greater than 15
(println "Filter (Greater than 15):"
         (filter #(> % 15) marks))

;; Reduce - Find maximum mark
(println "Reduce (Maximum):"
         (reduce max marks))

;; Take - Select first three elements
(println "Take 3:"
         (take 3 marks))

;; Drop - Remove first two elements
(println "Drop 2:"
         (drop 2 marks))

;; Threading macro
(println "\nThreading Macro (Marks):")

(def result2
  (->> marks
       (filter #(> % 15))
       (map inc)
       (drop 1)
       (take 3)))

(println "Filtered, Incremented and Sliced:" result2)

(println "\nSequence transformations completed successfully.")
