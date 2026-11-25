(defproject student-pipeline "0.1.0-SNAPSHOT"
  :description "Clojure Student Performance Data Pipeline"

  :dependencies [[org.clojure/clojure "1.12.0"]
                 [org.mongodb/mongodb-driver-sync "5.2.1"]
                 [ring/ring-core "1.13.0"]
                 [ring/ring-jetty-adapter "1.13.0"]
                 [cheshire "5.13.0"]
                 [org.slf4j/slf4j-simple "2.0.16"]]

  :main student-pipeline.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})
  