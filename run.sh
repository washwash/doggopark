rebuildApp() {
  docker-compose stop
  docker-compose build
  docker-compose up &

}

rebuildApp

#  vim creates 4913 file on almost every edit to check if a directory is writable
inotifywait -mr --exclude "(.*\.(swp|swo)|4913)" -e close_write back front | while read line
do 
  rebuildApp
done


