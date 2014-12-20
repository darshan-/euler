// Built with `R CMD SHLIB 012d.c`

int countFactors(int n) {
  int count = 0;

  for (int i = 1; ; i++) {
    if (n % i == 0) { // i is a factor
      count++;
      if (i * i == n) {
        return count * 2 - 1;
      } else if (i * i > n) {
        return (count - 1) * 2;
      }
    }

    if (i * i > n) {
      return count * 2;
    }
  }
}

void firstWithNFactors(int *n, int *triNum) {
  *triNum = 0;

  for (int i = 1; ; i++) {
    *triNum += i;
    if (countFactors(*triNum) > *n) { break; }
  }
}
