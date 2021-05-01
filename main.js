const fetch = require("node-fetch");
const cheerio = require("cheerio");
const express = require("express");
const fs = require("fs");
const path = require("path");

process.on("SIGINT", function () {
	process.exit();
});

const getData = async () => {
	const res = await fetch("https://www.expressvpn.com/what-is-my-ip");
	const html = await res.text();
	const $ = cheerio.load(html);

	const data = {
		ip: "",
		vpn: false,
		location: null,
	};

	try {
		const $ip = $(".ip-address > span");
		data.ip = $ip.text().trim();
		data.vpn = $ip.hasClass("green");
		data.location = $("h6:contains(Location)")
			.parent()
			.find("h4")
			.text()
			.trim();
	} catch (error) {}

	return data;
};

const app = express();

const template = fs.readFileSync(
	path.resolve(__dirname, "template.html"),
	"utf8",
);

app.get("/api", async (req, res) => {
	try {
		const data = await getData();
		res.send(data);
	} catch (error) {
		res.send(error);
	}
});

app.get("/", async (req, res) => {
	try {
		const data = await getData();

		const html = template
			.replace(/{{ bodyClass }}/gi, data.vpn ? "green" : "")
			.replace(
				/{{ text }}/gi,
				data.vpn ? "Connected to ExpressVPN" : "Not connected!",
			)
			.replace(/{{ ip }}/gi, data.ip)
			.replace(/{{ location }}/gi, data.location);

		res.contentType("html");
		res.send(html);
	} catch (error) {
		res.send(error);
	}
});

const port = process.env.PORT ?? 3000;
app.listen(port, () => {
	console.log("Listening on 0.0.0.0:" + port);
});
