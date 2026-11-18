(ns exp17)

(println "========================================")
(println "  CLOJURE IMMUTABLE DATA STRUCTURES")
(println "========================================")

;; 1. LIST OPERATIONS
(println "\n--- 1. LIST OPERATIONS ---")

(def original-list '(10 20 30 40))
(def updated-list (conj original-list 5))

(println "Original List:" original-list)
(println "First Element:" (first original-list))
(println "Remaining Elements:" (rest original-list))
(println "New List after conj:" updated-list)
(println "Original List after operation:" original-list)

;; 2. VECTOR OPERATIONS
(println "\n--- 2. VECTOR OPERATIONS ---")

(def original-vector [100 200 300 400])
(def updated-vector (conj original-vector 500))

(println "Original Vector:" original-vector)
(println "Element at index 2:" (get original-vector 2))
(println "Number of Elements:" (count original-vector))
(println "New Vector after conj:" updated-vector)
(println "Original Vector after operation:" original-vector)

;; 3. MAP OPERATIONS
(println "\n--- 3. MAP OPERATIONS ---")

(def original-map
  {:id 101 :name "Rahul" :department "CSE"})

(def added-map
  (assoc original-map :marks 85))

(def removed-map
  (dissoc original-map :department))

(println "Original Map:" original-map)
(println "Student Name:" (get original-map :name))
(println "Map after assoc:" added-map)
(println "Map after dissoc:" removed-map)
(println "Original Map after operations:" original-map)

;; 4. SET OPERATIONS
(println "\n--- 4. SET OPERATIONS ---")

(def original-set #{10 20 30 40})

(def added-set
  (conj original-set 50))

(def removed-set
  (disj original-set 20))

(println "Original Set:" original-set)
(println "Set after conj:" added-set)
(println "Set after disj:" removed-set)
(println "Contains 30?" (contains? original-set 30))
(println "Original Set after operations:" original-set)

(println "\n========================================")
(println "ALL FOUR DATA STRUCTURES DEMONSTRATED")
(println "Original collections remain unchanged.")
(println "========================================")