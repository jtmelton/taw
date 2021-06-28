# taw
A zero-dependency tool for finding all the file extensions and their count in a directory.

## Building
Maven is required for building.
```bash
# Building
CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo . 

# Running
./taw-inputDirectory /path/to/my/directory -outputFile myFileExtensions.json
```

## Arguments
Arg | Description | Required
------ | ------ | ------
inputDirectory | Directory containing a directory to analyze (Required) | true
outputFile | Output File | true

## Report

#### The output json is as follows:

```Javascript
  {
    "extension": ".go",
    "count": 74
  },
  {
    "extension": ".java",
    "count": 15
  },
  {
    "extension": ".docx",
    "count": 83
  },
```

## Docker Support
As taw is a zero-dependency golang app, the docker container is built from scratch, which makes it small and secure. 

```bash
docker build -t <tag_of_your_choice> .
```

Your docker container will need at least one mount point for the directory containing your app. Here is an example.
```bash
docker run --read-only -v /source/path/to/app:/path/to/app/in/container -v /source/path/to/resultsdir:/path/to/resultsdir/in/container -it <tag_built_with> -inputDirectory /path/to/app/in/container -outputFile /path/to/resultsdir/in/container
```
The --mount variant of mounting a volume can also be used if desired. If you want to write the output to a location outside of your container, then you will have to set a second mount point or re-use the existing one. If memory issues are encounterd, try running container with increased memory using the -m argument.
