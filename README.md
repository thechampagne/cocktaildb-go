# CocktailDB

[![](https://img.shields.io/github/v/tag/thechampagne/cocktaildb-go?label=version)](https://github.com/thechampagne/cocktaildb-go/releases/latest) [![](https://img.shields.io/github/license/thechampagne/cocktaildb-go)](https://github.com/thechampagne/cocktaildb-go/blob/main/LICENSE)

TheCocktailDB API client for **Go**.

### Download

```
go get github.com/thechampagne/cocktaildb-go
```

### Example

```go
func main() {
	cocktail, err := cocktaildb.Search("Queen Charlotte")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cocktail)
}
```

### License

CocktailDB is released under the [Apache License 2.0](https://github.com/thechampagne/cocktaildb-go/blob/main/LICENSE).

```
 Copyright 2022 XXIV

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
```