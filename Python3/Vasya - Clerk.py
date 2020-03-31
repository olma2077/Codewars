def tickets(people):
    price = 25
    cashbox = []
    for i in people:
        cashbox.append(i)
        if i - price > 0:
            change = i - price
            cashbox.sort(reverse=True)
            cashboxFreeze = cashbox.copy()
            for j in cashboxFreeze:
                if change - j == 0:
                    change -= j
                    cashbox.remove(j)
                    break
                elif change - j > 0:
                    change -= j
                    cashbox.remove(j)
            if change > 0:
                return "NO"
    return "YES"
