# go-aat

Go package for working with Getty Art & Architecture Thesaurus (AAT) XML data

## tools

```
$> make cli
go build -mod vendor -ldflags="-s -w" -o bin/aat2csv cmd/aat2csv/main.go
```

### aat2csv

Parse Getty AAT XML data and emit as CSV data to STDOUT with following columns: id, term, preferred, languages.

```
$> ./bin/aat2csv -h
Parse Getty AAT XML data and emit as CSV data to STDOUT with following columns: id, term, preferred, languages.
Usage:
	 ./bin/aat2csv [options]
  -terms string
    	The path to a local file on disk containing Getty AAT XML vocabulary data. If empty that data will be retrieved from the Getty servers.
```

For example:

```
./bin/aat2csv | less
preferred,id,term,languages
1,300000000,Top of the AAT hierarchies,70051/English
0,300000000,AAT Root,70051/English
0,300000000,Top van de AAT-hiërarchieën,70261/Dutch
0,300000000,藝術與建築詞典根目錄,72551/Chinese (traditional)
0,300000000,i shu yü chien chu tz'u tien ken mu lu,72582/Chinese (transliterated Wade-Giles)
0,300000000,yi shu yu jian zhu ci dian gen mu lu,72586/Chinese (transliterated Pinyin without tones)
0,300000000,yì shù yǔ jiàn zhú cí diǎn gēn mù lù,72584/Chinese (transliterated Hanyu Pinyin)
0,300000000,藝術和建築索引典之最高層級,72551/Chinese (traditional)
0,300000000,藝術與建築索引典最上層,72551/Chinese (traditional)
1,300189559,male,70051/English
0,300189559,masculino,70641/Spanish
0,300189559,mannelijk,70261/Dutch
0,300189559,雄性,72551/Chinese (traditional)
0,300189559,hsiung hsing,72582/Chinese (transliterated Wade-Giles)
0,300189559,xiong xing,72586/Chinese (transliterated Pinyin without tones)
0,300189559,xióng xìng,72584/Chinese (transliterated Hanyu Pinyin)
1,300189557,female,70051/English
... and so on
```

## Notes

This package does not use a streaming XML parser so the entirety of the AAT XML data is read in to memory before parsing occurs.

## See also

* http://aatdownloads.getty.edu/