package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(name string, pageURL *string) (RespEncounters, error) {
	url := baseURL + "/location-area/" + name
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		encounterResp := RespEncounters{}
		err := json.Unmarshal(val, &encounterResp)
		if err != nil {
			return RespEncounters{}, err
		}
		return encounterResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespEncounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespEncounters{}, err
	}

	encounterResp := RespEncounters{}
	err = json.Unmarshal(dat, &encounterResp)
	if err != nil {
		return RespEncounters{}, err
	}
	c.cache.Add(url, dat)
	return encounterResp, nil
 }
