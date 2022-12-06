#!/bin/sh
EX=$1
mkdir $EX
cp template/main.go $EX/$EX.go
cp template/main_test.go $EX/"$EX"_test.go
cp template/*.txt $EX/
sed "s/EX/$EX/g" template/run.sh > $EX/run.sh
chmod +x $EX/run.sh
cd $EX
go mod init "day$EX"
go work init ./
