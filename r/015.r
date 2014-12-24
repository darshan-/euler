#!/usr/bin/Rscript

DIMEN <- 20

PD <- DIMEN + 1 ## Point dimensions; a 2x2 grid has 3x3 points, and so on

m <- matrix(nrow=PD, ncol=PD)

m[1,] <- m[,1] <- 1

for (r in 2:PD) {
    for (c in 2:PD) {
        m[r,c] <- sum(m[r,c-1], m[r-1,c])
    }
}

cat(m[PD,PD], "\n")
