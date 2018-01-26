import yaml, json, argparse

parser = argparse.ArgumentParser()
parser.add_argument('--inputFile','-i', required=True, help='path to the input file')
parser.add_argument('--outputFile','-o', required=True, help='path to the output file')
args = parser.parse_args()

def loadYaml(path):
	file = open(path, 'r')
	return yaml.load(file)

def convertToJson(yamlData):
	return json.dumps(yamlData)

def writeToDisk(outputFile, jsonData):
	file = open(outputFile, 'w')
	file.write(jsonData)
	file.close

print ("Loading yaml")
yaml = loadYaml(args.inputFile)
print ("Converting to JSON")
jsonData = convertToJson(yaml)
print ("Writing JSON to dsik")
writeToDisk(args.outputFile, jsonData)
