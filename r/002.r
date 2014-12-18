LIMIT <- 4*1000*1000

sum <- 0
a   <- 1
b   <- 2

while (b < LIMIT) {
    tmp <- a
    a   <- b
    b   <- tmp + b

    # On the last iteration, b will be greater than LIMIT and a will be the
    #  largest fib under LIMIT.  On the first iteration, a becomes 2 and gets
    #  added to sum.  If a's initial value were even, we would need to add it
    #  to sum before the loop.
    if (a %% 2 == 0) {
        sum <- sum + a
    }
}

print(sum)
