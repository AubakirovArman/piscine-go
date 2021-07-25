touch -a -m -t 198601050000.00 0
touch -a -m -t 198611130001.00 1
touch -a -m -t 198803050010.00 2
touch -h -m -t 199002160011.00 3
touch -a -m -t 199010070200.00 4
touch -a -m -t 199011070101.00 5
touch -a -m -t 199102080110.00 6
touch -a -m -t 199103080111.00 7
touch -a -m -t 199405201100.00 8
touch -a -m -t 199406101101.00 9
touch -a -m -t 199504101110.00 A
ls -l --time-style='+%F %R' | sed 1d 
sudo tar --totals --create --verbose --file done.tar 0 1 2 3 4 5 6 7 8 9 A
git add done.tar
git commit -m "done2"
git push 




