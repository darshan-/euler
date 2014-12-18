#!/usr/bin/Rscript

nextPrime <- (function() {
    cur <- 1
    primes <- NULL

    isNextPrime <- function(n) {
        for (prime in primes) {
            if (n %% prime == 0) {
                return(FALSE)
            }
        }
        TRUE
    }

    function() {
        cur <<- cur + 1
        while (!isNextPrime(cur)) {
            cur <<- cur + 1
        }

        primes <<- append(primes, cur)

        cur
    }
})()

maxPrimeFactor <- function(n) {
    factors <- NULL

    while(TRUE) {
        i <- nextPrime()
        while (n %% i == 0) {
            factors <- append(factors, i)
            cat(" +++ Factors:", factors, "\n")
            if (n == i) { return(i) }
            n <- n / i
            cat(" * New n:", n, "\n")
        }
    }
}

cat(maxPrimeFactor(600851475143), "\n")
