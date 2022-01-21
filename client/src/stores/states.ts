import { writable } from 'svelte/store';
import { States } from '../utils/constants';

export const state = writable(States.result)