(ns exp18)

(println "========================================")
(println "       SEQUENCE TRANSFORMATIONS")
(println "========================================")

;; Function to demonstrate sequence operations
(defn process-sequence [label collection]

  (println "\n----------------------------------------")
  (println label)
  (println "----------------------------------------")

  (println "Original Collection:" collection)

  ;; 1. MAP - Square every element
  (let [squared (map #(* % %) collection)]

    (println "After map (squares):" squared)

    ;; 2. FILTER - Keep only even numbers
    (let [even-numbers (filter even? squared)]

      (println "After filter (even squares):" even-numbers)

      ;; 3. REDUCE - Calculate sum
      (let [total (reduce + 0 even-numbers)]

        (println "After reduce (sum):" total))))

  ;; 4. TAKE - Select first three elements
  (println "After take (first 3):"
           (take 3 collection))

  ;; DROP - Skip first two elements
  (println "After drop (skip first 2):"
           (drop 2 collection))

  ;; THREADING MACRO ->>
  ;; Chain map, filter, take and reduce
  (let [result
        (->> collection
             (map #(* % %))
             (filter even?)
             (take 3)
             (reduce + 0))]

    (println "Threading Macro Result:" result))

  (println "Original Collection Remains:" collection))


;; TEST 1: First collection (vector)
(def numbers-one [1 2 3 4 5 6 7 8])

(process-sequence "TEST 1 - VECTOR COLLECTION"
                  numbers-one)


;; TEST 2: Second collection (list)
(def numbers-two '(2 4 6 8 10 12 14))

(process-sequence "TEST 2 - LIST COLLECTION"
                  numbers-two)


;; Demonstrate -> threading macro
(println "\n--- THREADING MACRO (->) ---")

(def value 5)

(def transformed-value
  (-> value
      (* 2)
      (+ 10)
      (- 3)))

(println "Original Value:" value)
(println "Transformed Value:" transformed-value)

(println "\n========================================")
(println "SEQUENCE TRANSFORMATIONS COMPLETED")
(println "Both collections tested successfully.")
(println "========================================")