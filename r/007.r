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

for (i in 1:1e4) {
    nextPrime()
}

cat(nextPrime(), "\n")
