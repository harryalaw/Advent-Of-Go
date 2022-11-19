YEAR=$ADVENT_YEAR
USER_DAY=$1

if [[ ${#USER_DAY} == 1 ]]; then
  echo $USER_DAY
  echo ${#USER_DAY}
  DAY="0$USER_DAY"
elif [[ ${#USER_DAY} == 2 ]]; then
  DAY=$USER_DAY
else 
  echo "Day should be between 1 and 25"
  return
fi

if [[ -n $2 ]]; then
  YEAR=$2
fi

if [[ -z $YEAR ]]; then
  echo "No year provided"
  echo "Please set the "ADVENT_YEAR" env variable"
  echo "Or provide a year argument"
  return
fi

# Get advent of code data
if [[ -z $ADVENT_COOKIE ]]; then
  echo "Please provide cookie for advent of code"
  return
fi

RES=$(curl -s https://adventofcode.com/$YEAR/day/$USER_DAY/input --cookie session=$ADVENT_COOKIE)

ERROR_TEXT="Puzzle inputs differ by user. Please log in to get your puzzle input."

if [[ "$ERROR_TEXT" == "$RES" ]]; then
  echo "$ERROR_TEXT==$RES"
  echo "Auth failed, please update ADVENT_COOKIE env variable"
  echo $RES
  return
fi

FILE_DIR="$YEAR/$DAY"
mkdir -p $FILE_DIR

echo $RES > $FILE_DIR/data.txt

# Initialize templates
cp -n templates/* $FILE_DIR/
