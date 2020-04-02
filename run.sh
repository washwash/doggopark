rebuildApp() {
  docker-compose down
  docker-compose build
  docker-compose up &

}

rebuildApp

inotifywait -m -r -e close_write back front | while read line
do
  rebuildApp
done


