import { writable } from "svelte/store"

export const word = writable(0)
export const char = writable(0)
export const wordsPerMinute = writable(0)
export const incorrectChars = writable(0)
export const totalChars = writable(0)