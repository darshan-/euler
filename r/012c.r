#!/usr/bin/Rscript

library(Rcpp)

cppFunction('int countFactors(int n) {
    int count = 0;

    for (int i = 1; ; i++) {
        if (n % i == 0) { // i is a factor
            count++;
            if (i * i == n) {
                return count * 2 - 1;
            } else if (i * i > n) {
                return (count - 1) * 2;
            }
        }

        if (i * i > n) {
            return count * 2;
        }
    }
}')

GOAL <- 500

triNum <- 0
i <- 1

repeat {
    triNum <- triNum + i
    if (countFactors(triNum) > GOAL) { break }
    i <- i + 1
}

cat(triNum, "\n")
