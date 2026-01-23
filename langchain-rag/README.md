# LangChain exploratory notes

## Summary

Brief findings from exploratory LangChain usage focusing on modern ChatModel patterns and RAG options.

## LangChain Insights

- ChatModels (e.g., ChatOllama) are modern alternatives to legacy LLM classes. They support multi-modality (text, images, etc.) & return structured objects (messages, roles, content, attachments, metadata) rather than plain strings.
- Given a goal; An Agent = Loop + LLM + State + Tools(optional) until goal is evaluated or stop condition is reached.
- Tools require manual invocation when using `Chat*` models, the LLM produces an output in its `tool_calls` property, you must loop over it and invoke the tool and then pass its return value back to the LLM. This is not true when using `createAgent` api as it internally invokes the tools requested by LLM.
- If you simply use the `Chat*` model, you must manage all of the ceremony, i.e., `sessionMemory`, `toolInvocations` etc manually in your code. `createAgent` internally handles that - provided you pass it a `checkPointer` instance - giving you a clean API.

## RAG implementation options

- Our current implementation is manually invoking the `retrieval` tool regardless the user's query. The user may have said `hi` but our system will still retrieve the relevant chunks from the db. (Naive/AlwaysRetrieve)
- Another implementation is to `bindTools` to the model directly. It'll allow the model to decide whether to call the tool or not. This will overcome the issue mentioned earlier. (Agentic)
- Yet another approach is to utilize multiple models, one model will always be invoked to determine or _classify_ the user's query. If it requires retrieval, it will invoke the tool. If not, it'll simply let the RAG_MODEL handle it. (Router)
