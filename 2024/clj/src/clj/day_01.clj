(ns day-01
  (:require
   [clojure.java.io :as io]
   [clojure.string :as str]))

(def input (slurp (io/resource "day_01.txt")))

(defn part-01 [input]
  (let [pairs (parse input)]
       (reduce + (distances pairs))))

(defn parse [input]
  (->> input
       (#(str/split % #"\n"))
       (map (fn [line] (map
                        #(Integer. %)
                        (str/split line #"\s+"))))
       ((fn [pairs]
          [(-> (mapv first pairs) sort)
           (-> (mapv second pairs) sort)]))
       (apply map vector)))

(defn distances [pairs]
  (map (fn [[left right]]
         (if (> left right)
           (- left right)
           (- right left))) pairs))
