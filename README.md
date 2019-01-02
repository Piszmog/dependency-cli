# Dependency CLI
Microservice applications are becoming more and more widely used. As the number of microservice applications 
increase, it becomes harder and harder to manage all the dependencies. Microservices may depend on parent projects, 
and a number of libraries. And these parents can also depend on a number of parents and libraries. 

Updating all the dependencies can become time consuming and introduce risk. This CLI tool is meant to make updating 
dependencies of projects easier and faster in an automated fashion.

## Configuration File
The tool requires a configuration file in JSON format to run.

```json
{
  "updateDependencies": true,
  "updateParent": true,
  "includes": [
    {
      "groupId": "com.fasterxml.jackson.dataformat",
      "artifactId": "*"
    },
    ...
  ],
  "excludes": [
    {
      "groupId": "org.apache.commons",
      "artifactId": "commons-lang3"
    },
    ...
  ],
  "mavenProjects": [
    {
      "baseDirectory": "${path to directory contining the below list of projects}",
      "projects": [
        "${project 1}",
        "${project 2}",
        "${project 3}",
        ...
      ]
    },
    ...
  ]
}
```

Where,

| Field | Type | Description |
| --- | --- | --- |
| `updateDependencies` | boolean | Determines if the dependencies of a project should be updated to he latest released versions |
| `updateParent` | boolean | Determines if the parent of a project should be updated to he latest released version |
| `includes` | list | List of dependencies determined by `groupId` and `artificatId` that will only be updated - wildcard `*` can be used |
| `excludes` | list | List of dependencies determined by `groupId` and `artificatId` that will __NOT__ be updated - wildcard `*` can be used |
| `mavenProjects` | list | List of projects that will be updated based on the above configurations |

## Running the CLI
First, download the [latest artifact](https://github.com/Piszmog/dependency-cli/releases) of the CLI available.

Then run with a `-f` flag and provide the path to the configuration file.

### Windows
`dependency-cli-1.0.0.exe -f /path/to/configuration/file/configuration.json`

### Linux
`./dependency-cli-1.0.0 -f /path/to/configuration/file/configuration.json`

## Future Implementations
* SNAPSHOT support
  * Currently, only latest release versions are used
* GIT support
  * Ability to commit and push dependency changes to GIT branch
* Release
  * Run a release command after changes have been performed
* CI/CD support
  * Execute Bamboo/Jenkins/ect... jobs after changes have been performed
  
