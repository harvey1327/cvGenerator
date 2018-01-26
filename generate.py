import yaml, json, subprocess

def loadYaml(path):
	file = open(path, 'r')
	return yaml.load(file)

def convertToJson(yamlData):
	return json.dumps(yamlData)

def writeToDisk(outputFile, jsonData):
	file = open(outputFile, 'w')
	file.write(jsonData)
	file.close

def dockerRun():
	subprocess.call("")

if __name__ == '__main__':
	print ("Loading yaml")
	yaml = loadYaml('./resume/resume.yml')
	print ("Converting to JSON")
	jsonData = convertToJson(yaml)
	print ("Writing JSON to dsik")
	writeToDisk('./resume/resume.json', jsonData)
#	print ("Running docker image")
#	dockerRun()
