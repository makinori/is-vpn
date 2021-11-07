import { Application, Router } from "./deps.ts";
import { services } from "./services.ts";

const SERVICE = Deno.env.get("SERVICE");
if (SERVICE == null) {
	console.error("Please provide SERVICE environment variable");
	Deno.exit(1);
}

if (services[SERVICE] == null) {
	console.error(`Service "${SERVICE}" not supported`);
	Deno.exit(1);
}

const dirname = new URL(".", import.meta.url).pathname;
const template = await Deno.readTextFile(dirname + "template.html");

const router = new Router();

router.get("/api", async ctx => {
	try {
		ctx.response.body = await services[SERVICE]();
	} catch (error) {
		ctx.response.status = 500;
		ctx.response.body = { error: error.toString() };
	}
});

router.get("/", async ctx => {
	try {
		const status = await services[SERVICE]();
		ctx.response.body = template
			.replace(
				/{{ text }}/g,
				status.status
					? "Connected to " + status.name
					: "Not connected to " + status.name + "!",
			)
			.replace(/{{ ip }}/g, status.ip)
			.replace(/{{ location }}/g, status.location)
			.replace(/{{ bodyClass }}/g, status.status ? "green" : "");
	} catch (error) {
		ctx.response.status = 500;
		ctx.response.body = { error: error.toString() };
	}
});

const app = new Application();
app.use(router.routes());
app.use(router.allowedMethods());
app.use(ctx => {
	ctx.response.status = 404;
	ctx.response.body = { error: "Not found" };
});

const port = 8080;
console.log("Listening on http://127.0.0.1:" + port);
await app.listen({ port });
