import { writable } from "svelte/store";

export const selectedMenu = writable(0);
export const selectedWallet = writable("");
export const selectedClient = writable("");
