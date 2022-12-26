package area

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
)


type Species struct {
  Id          string   `json:"id"`
  Name        string   `json:"name"`
  Popularity  int      `json:"popularity"`
  Location    string   `json:"location"`
  Availability []string `json:"availability"`
  Residents   []struct {
    Name string `json:"name"`
    Sex  string `json:"sex"`
    Age  int    `json:"age"`
  } `json:"residents"`
}

type ZooData struct {
  Species []Species `json:"species"`
}

func GetSpeciesByIds(ids ...string) ([]Species, error) {
  // ler o arquivo JSON e armazenar os dados em uma variável do tipo ZooData
  data, err := ioutil.ReadFile("./data/zoo_data.json")
  if err != nil {
    return nil, err
  }
  
  var dataMap ZooData
  json.Unmarshal(data, &dataMap)

  // criar o array de retorno da função
  var result []Species
	

  // adicionar as espécies correspondentes aos ids passados como parâmetro
  for _, id := range ids {
    found := false
    for _, species := range dataMap.Species {
      if species.Id == id {
        found = true
        result = append(result, species)
        break
      }
    }
    if !found {
      return nil, fmt.Errorf("species with id %s not found", id)
    }
  }
  return result, nil
}
