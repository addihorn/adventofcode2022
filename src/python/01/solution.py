
f = open("input.txt", "r")

elveList = f.read().split("\n\n")
print (len(elveList))
dictElveCaloriesCarried = {}

for elveNumber in range(len(elveList)):
    elveListItems = elveList[elveNumber].split("\n")
    elveCaloriesCarried = sum([int(x) for x in elveListItems])
    dictElveCaloriesCarried[elveNumber] = elveCaloriesCarried




print("Elve with max Calories: " + str(max(dictElveCaloriesCarried, key=dictElveCaloriesCarried.get)+1))
print("max Calories: " +   str(max(dictElveCaloriesCarried.values())))


caloriesList = list(dictElveCaloriesCarried.values())
caloriesList.sort(reverse=True)

caloriesByTopX = caloriesList
caloriesByTopX
for x in range(1, len(caloriesList)):
    caloriesByTopX[x] = caloriesByTopX[x-1] + caloriesList[x]

print("Calories carried by top 3 elves: " + str(caloriesByTopX[3-1]))