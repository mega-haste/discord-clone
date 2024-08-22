export function setCookie(cname: string, cvalue: string, exdays: number) {
	const d = new Date();
	d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
	let expires = "expires="+d.toUTCString();
	document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}


export function getCookie(cname: string): string | undefined {
	let name = cname + "=";
	let ca = document.cookie.split(';');
	for(let c of ca) {
		c = c.trimStart();
		if (c.indexOf(name) == 0)
			return c.substring(name.length, c.length);
	}
	return undefined;
}

export function checkCookie(cname: string): boolean {
	return Boolean(getCookie(cname));
}


export function getAllCookies(): string {
	return document.cookie;
}
