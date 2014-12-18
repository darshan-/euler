collect <- function(dict, step, limit){
    i <- step
    while (i < limit) {
        n <- i
        dict[i] <- n
        i <- i + step
    }

    dict
}

LIMIT <- 1000

d <- c()

d <- collect(d, 3, LIMIT)
d <- collect(d, 5, LIMIT)

print(sum(na.omit((d))))
