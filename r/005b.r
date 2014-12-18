#!/usr/bin/Rscript

MAX <- 20

primesToMax <- function(max) {
    primes <- NULL
    addNextPrime <- (function() {
        cur <- 1

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

    n <- 0
    while (n < max) {
        n <- addNextPrime()
    }

    if (primes[length(primes)] > max) {
        primes <- primes[1:length(primes)-1]
    }

    primes
}

primes <- primesToMax(MAX)

factors <- function(n) {
    fs <- NULL

    for (prime in primes) {
        while (n %% prime == 0) {
            fs <- append(fs, prime)
            if (n == prime) { break }
            n <- n / prime
        }
    }

    fs
}

# Returns the number of x in v
count <- function(v, x) {
    length(v[v == x])
}

# Returns a vector of n x
vofx <- function(x, n = 1) {
    v <- NULL
    v[1:n] <- x
    v
}

mergef <- function(v1, v2) {
    fs <- NULL

    for (f in unique(c(v1, v2))) {
        n <- max(count(v1, f), count(v2, f))
        fs <- append(fs, vofx(f, n))
    }

    fs
}

allf <- NULL
for (n in 1:MAX) {
    allf <- mergef(allf, factors(n))
}

cat(prod(allf), "\n")
