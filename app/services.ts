import { DOMParser } from "./deps.ts";

interface VpnStatus {
	ip: string;
	status: boolean;
	location: string;
	name: string;
}

async function mullvad() {
	const res = await fetch("https://ipv4.am.i.mullvad.net/json");
	const { ip, country, city, mullvad_exit_ip, mullvad_exit_ip_hostname } =
		await res.json();
	return {
		ip,
		status: mullvad_exit_ip,
		location: `${country}, ${city} (${mullvad_exit_ip_hostname})`,
		name: "Mullvad",
	};
}

async function nordvpn() {
	const res = await fetch(
		"https://nordvpn.com/wp-admin/admin-ajax.php?action=get_user_info_data",
	);
	const { ip, status, location } = await res.json();
	return {
		ip,
		status,
		location,
		name: "NordVPN",
	};
}

async function expressvpn() {
	const res = await fetch("https://www.expressvpn.com/what-is-my-ip");
	const doc = new DOMParser().parseFromString(await res.text(), "text/html")!;

	const ipEl = doc.querySelector(".ip-address > span");

	const ip = ipEl?.textContent.trim() ?? "Unknown";
	const status = ipEl?.className.includes("green") ?? false;
	const location =
		Array.from(doc.querySelectorAll("h6"))
			.find(h => h.textContent == "Location")
			?.parentElement?.querySelector("h4")
			?.textContent.trim() ?? "Unknown";

	return {
		ip,
		status,
		location,
		name: "ExpressVPN",
	};
}

export const services: { [service: string]: () => Promise<VpnStatus> } = {
	mullvad,
	nordvpn,
	expressvpn,
};
