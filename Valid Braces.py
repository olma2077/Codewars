def validBraces(string):
    buffer = []
    for i in string:
        if i == '(' or i == '[' or i == '{':
            buffer.append(i)
        else:
            try:
                a = buffer.pop()
            except IndexError:
                return False
            if i == ')' and a != '(':
                return False
            elif i == ']' and a != '[':
                return False
            elif i == '}' and a != '{':
                return False
    if buffer:
        return False
    else:
        return True
        
