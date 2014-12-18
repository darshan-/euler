def countFactors(n)
  count = 0

  i = 0
  while (i += 1)
    if n % i == 0 # i is a factor
      count += 1
      if i * i == n
        return count * 2 - 1
      elsif i * i > n
        return (count - 1) * 2
      end
    end

    if i * i > n
      return count * 2
    end
  end
end

goal = 500

triNum = 0

i = 0
while (i += 1)
  triNum += i
  if countFactors(triNum) > goal
    break
  end
end

puts triNum
