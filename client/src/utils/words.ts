import axios from "axios"
import variables from "../lib/variables"

export function getWords(): string[] {
	return [
		"rack",
		"taxi",
		"dish",
		"microphone",
		"sharp",
		"make",
		"talented",
		"solo",
		"race",
		"multimedia",
		"red",
		"proclaim",
		"name",
		"tycoon",
		"evolution",
		"effect",
		"linger",
		"learn",
		"lighter",
		"redundancy",
		"drift",
		"lesson",
		"cutting",
		"section",
		"house",
		"feminine",
		"future",
		"census",
		"exceed",
		"bathtub",
		"topple",
		"center",
		"world",
		"overeat",
		"surprise",
		"posture",
		"verdict",
		"leg",
		"gossip",
		"opposite",
		"operation",
		"goal",
		"traction",
		"recruit",
		"shelter",
		"reflect",
		"banner",
		"loss",
		"introduce",
		"dry",
	]
}

export const getRandomWords = async () => {
	console.log(`${variables.baseURL}/randomTopWords`)
	const res = await axios.get(`http://localhost:8000/randomTopWords`)
	const words = await res.data.words

	if (res.status === 200) {
		return words
	} else {
		throw new Error(words)
	}
}