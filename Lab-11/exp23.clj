(defn divide [dividend divisor]
  (quot dividend divisor))

(defn string-length [^String value]
  (.length value))

(defn validate-quantity [quantity]
  (if (neg? quantity)
    (throw (ex-info "Quantity cannot be negative"
                    {:quantity quantity
                     :fallback 0}))
    quantity))

(defn run-case [label operation]
  (println "\nCase:" label)
  (try
    (operation)
    (catch ArithmeticException _
      (println "Division by zero. Using 0.")
      0)
    (catch NullPointerException _
      (println "Missing string. Using length 0.")
      0)
    (catch clojure.lang.ExceptionInfo exception
      (let [{:keys [quantity fallback]} (ex-data exception)]
        (println (ex-message exception))
        (println "Invalid quantity:" quantity)
        (println "Using fallback quantity:" fallback)
        fallback))
    (finally
      (println "Finished:" label))))

(println "Result:" (run-case "Division by zero" #(divide 10 0)))
(println "Result:" (run-case "Missing string" #(string-length nil)))
(println "Result:" (run-case "Invalid quantity" #(validate-quantity -3)))
(println "Result:" (run-case "Valid quantity" #(validate-quantity 5)))