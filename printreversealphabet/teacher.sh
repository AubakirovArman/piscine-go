number_interview=$(head -n 179 streets/Buckingham_Place | tail -n 1 | awk -F# '{print $2}')
echo $number_interview
cat interviews/interview-$number_interview
color=$(cat interviews/interview-$number_interview | grep "Describes" | awk -F. '{print $2}' |awk -F, '{print $1}' | awk '{print $5}')
mark=$(cat interviews/interview-$number_interview | grep "Describes" | awk -F. '{print $2}' |awk -F, '{print $1}' | awk '{print $6}')    
number_car_first=$(cat interviews/interview-$number_interview | grep "Describes" | awk -F. '{print $2}' |awk -F, '{print $2}' | awk -F\" '{print $2}')
number_car_last=$(cat interviews/interview-$number_interview | grep "Describes" | awk -F, '{print $3}' | awk -F\" '{print $4}' )

echo "$(grep "$number_car_first" vehicles -A 10  | grep $mark  -A 3 -B 1 | grep -i "$color" -A 3  -B 2 | grep "Height: 6" -B 4)"
list="$(echo "$(grep "$number_car_first" vehicles -A 10  | grep $mark  -A 3 -B 1 | grep -i "$color" -A 3  -B 2 | grep "Height: 6" -B 4)")"

cat memberships/AAA memberships/Delta_SkyMiles memberships/Museum_of_Bash_History memberships/Terminal_City_Library| grep "$MAIN_SUSPECT" | wc -l