#!/usr/bin/Rscript

cur <- 1
#primes <- rep(NA, 10001)
primes <- vector(length=10001)
#primev <- rep(NA, 10001)
#primes <- na.omit(primev)

isNextPrime <- function(n) {
    for (prime in primes) {
        if (!prime) { return(TRUE) }
        if (n %% prime == 0) {
            return(FALSE)
        }
    }
    TRUE
}


for (i in 1:10001) {
    cur <- cur + 1
    while (!isNextPrime(cur)) {
        cur <- cur + 1
    }

    primes[i] <- cur
    #primes <- na.omit(primev)
}

cat(primes[10001], "\n")
