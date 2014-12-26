#!/usr/bin/Rscript

sumDivisors <- function(n) {
    sum <- 0
    i <- 1

    while (i^2 <= n) {
        if (n %% i == 0) { # i is a factor
            sum <- sum + i
            if (i * i == n) {
                break
            }
            sum <- sum + (n / i)
        }

        i <- i + 1
    }

    return(sum - n)
}

sum <- 0

for (i in 1:9999) {
    n <- sumDivisors(i)
    if (n != i && sumDivisors(n) == i) {
        sum <- sum + i
    }
}

cat(sum, "\n")
