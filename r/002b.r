LIMIT <- 4*1000*1000
sum <- 0

nextfib <- (function() {
    a <- 0
    b <- 1

    function() {
        tmp <- a
        a   <<- b
        b   <<- tmp + b

        b
    }
})()

while (TRUE) {
    f <- nextfib()

    if (f > LIMIT) {
        break
    }

    if (f %% 2 == 0) {
        sum <- sum + f
    }
}

print(sum)
