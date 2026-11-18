(ns experiment23)

(println "====================================")
(println "     EXCEPTION HANDLING IN CLOJURE")
(println "====================================")


;; 1. ArithmeticException

(println "\n--- Test 1: Arithmetic Exception ---")

(try
  (println "Performing division...")

  (def result (quot 20 0))

  (println "Result:" result)

  (catch ArithmeticException e
    (println "Arithmetic Error:" (.getMessage e))
    (println "Recovered from division error."))

  (finally
    (println "Arithmetic test completed.")))


;; 2. NullPointerException

(println "\n--- Test 2: Null Pointer Exception ---")

(try
  (println "Checking string length...")

  (def text-value nil)

  (.length ^String text-value)

  (catch NullPointerException e
    (println "Null Pointer Error:" (.getMessage e))
    (println "Recovered from null value."))

  (finally
    (println "Null pointer test completed.")))


;; 3. Custom Exception using ex-info

(println "\n--- Test 3: Custom Exception ---")

(defn check-quantity [quantity]

  (when (neg? quantity)
    (throw
      (ex-info
        "Quantity cannot be negative."
        {:quantity quantity
         :error-type :invalid-quantity})))

  (println "Valid Quantity:" quantity))


(try
  (check-quantity -5)

  (catch clojure.lang.ExceptionInfo e

    (println "Custom Exception:" (ex-message e))

    (println "Additional Error Data:" (ex-data e))

    (println "Recovered from custom exception."))

  (finally
    (println "Custom exception test completed.")))


;; 4. Testing a valid input after errors

(println "\n--- Recovery Test ---")

(try
  (check-quantity 15)

  (catch clojure.lang.ExceptionInfo e
    (println "Error:" (ex-message e)))

  (finally
    (println "Recovery test completed.")))


(println "\n====================================")
(println " ALL EXCEPTION TESTS COMPLETED")
(println "====================================")