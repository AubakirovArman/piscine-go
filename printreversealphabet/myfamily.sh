
curl -s https://01.alem.school/assets/superhero/all.json | jq '.[]? | select(.id=='$(jq -n 'env.HERO_ID' | tr -d '\"')').connections.relatives' | tr -d '\"'