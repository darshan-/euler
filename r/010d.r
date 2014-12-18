#!/usr/bin/Rscript

LIMIT <<- 2e6
SQRT <<- floor(sqrt(LIMIT))
nums <<- vector(length=LIMIT)
sum <- 0

for (num in 2:LIMIT) {
    if (nums[num]) { next }

    sum <- sum + num

    if (num <= SQRT) {
        nums[seq(num^2, LIMIT, num)] <- TRUE
    }
}

cat(sum, "\n")
