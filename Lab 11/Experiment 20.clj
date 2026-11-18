(ns experiment20
  (:import [java.util ArrayList HashMap]))

(println "====================================")
(println "       JAVA INTEROPERABILITY")
(println "====================================")

;; 1. Calling Java static methods

(println "\n--- Java Static Methods ---")

(def number -25)
(def absolute-value (Math/abs number))

(println "Original Number:" number)
(println "Absolute Value:" absolute-value)

(def current-time (System/currentTimeMillis))

(println "Current Time in Milliseconds:" current-time)


;; 2. Creating Java ArrayList object

(println "\n--- Java ArrayList Operations ---")

(def java-list (ArrayList.))

(.add java-list "Apple")
(.add java-list "Banana")
(.add java-list "Mango")

(println "Java ArrayList:" java-list)

(println "First Element:" (.get java-list 0))
(println "List Size:" (.size java-list))

(.add java-list "Orange")

(println "Updated ArrayList:" java-list)


;; 3. Creating Java HashMap object

(println "\n--- Java HashMap Operations ---")

(def java-map (HashMap.))

(.put java-map "Name" "Abhinav")
(.put java-map "Course" "CSE")
(.put java-map "Semester" "7")

(println "Java HashMap:" java-map)

(println "Student Name:" (.get java-map "Name"))
(println "Course:" (.get java-map "Course"))

(.put java-map "K.R. Mangalam" "University")

(println "Updated HashMap:" java-map)


;; 4. Converting Clojure collections to Java

(println "\n--- Collection Conversion ---")

(def clojure-list [10 20 30 40 50])

(def converted-java-list
  (ArrayList. clojure-list))

(println "Original Clojure Collection:" clojure-list)
(println "Converted Java Collection:" converted-java-list)


;; 5. Converting Java collections to Clojure

(def converted-clojure-list
  (vec converted-java-list))

(println "Converted Back to Clojure:" converted-clojure-list)

(def clojure-map
  (into {} java-map))

(println "Java Map Converted to Clojure Map:" clojure-map)

(println "\n====================================")
(println " JAVA INTEROPERABILITY COMPLETED")
(println "====================================")