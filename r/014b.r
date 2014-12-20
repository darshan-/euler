#!/usr/bin/Rscript

GOAL<- 1e6
collatz <- vector(length=GOAL)
max   <- 1
which <- 1

for (start in 1:GOAL) {
    count <- 1
    n <- start

    while (n > 1) {
        if (n <= GOAL && collatz[n]) {
            count <- count + collatz[n] - 1 # Subtract one to not double-count n
            break
        }

        count <- count + 1

        if (n %% 2 == 0) {
            n <- n / 2
        } else {
            n <- 3 * n + 1
        }
    }

    collatz[start] <- count

    if (count > max) {
        max <- count
        which <- start
    }
}

cat(which, "\n")
