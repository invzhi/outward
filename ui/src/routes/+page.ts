// import { UserService } from "$lib/proto/outward/v1/user_service_connect";
// import { createPromiseClient } from "@connectrpc/connect";
// import { createConnectTransport } from "@connectrpc/connect-web";
// import type { PageLoad } from "./$types";

// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
export const prerender = true;

// export const load: PageLoad = async ({ fetch }) => {
//     const transport = createConnectTransport({
//       baseUrl: "http://localhost:8080",
//       // We pass `fetch` provided by Svelte to the Transport. The function
//       // behaves the same as native fetch(), but it inherits cookies, and
//       // it can make relative requests, so you don't have to specify an
//       // absolute baseUrl.
//       // For more information, see https://kit.svelte.dev/docs/load#making-fetch-requests
//       fetch,
//     });
  
//     const client = createPromiseClient(UserService, transport);
  
//     const request = {
//       workspaceId: "1434",
//     };
  
//     const response = await client.getUserList(request);
  
//     return {
//       request,
//       response,
//     };
//   };
