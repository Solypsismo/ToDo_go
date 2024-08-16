<script>
    import { removeObj, updateStatusToDone } from '$lib/utility';

    export let data;
    let taskTitle = "";

    const aggiungi = async (event) => {
        event.preventDefault();
        const form = event.target;

        const obj = Object.fromEntries(new FormData(form));
        console.log("aggiunta: ", obj)
        const resJSON = await fetch("http://localhost:8080/api/tasks", {
            method: "POST",
            body: JSON.stringify(obj),
        });

        if (resJSON.ok) {
            taskTitle = "";
            const { id } = await resJSON.json();
            tasks = [...tasks, { ...obj, id }];
            console.log(tasks);
        }
    };

    const rimuovi = async (event) => {
        const div = event.target.closest("div[data-task-id]");
        const id = div.dataset.taskId;

        let obj = JSON.stringify({ id })
        console.log("rimozione: ", obj)

        const resJSON = await fetch("http://localhost:8080/api/tasks", {
            method: "DELETE",
            body: obj,
        });

        if (resJSON.ok) {
            let lis = removeObj(tasks, id);
            tasks = lis;
        } else {
            alert("Undefined error");
        }
    };

    const aggiorna = async (event) => {
        const div = event.target.closest("div[data-task-id]");
        const id = div.dataset.taskId;

        let obj = JSON.stringify({ id })
        console.log("done: ", obj)

        const resJSON = await fetch("http://localhost:8080/api/tasks", {
            method: "PUT",
            body: obj,
        });

        if (resJSON.ok) {
            let ris = updateStatusToDone(tasks, id);
            tasks = ris;
        } else {
            alert("Undefined error");
        }
    };

    $: tasks = data.tasks;
</script>

<div
    class="w-screen h-screen bg-slate-300 flex justify-center items-center gap-10"
>
    <div class="w-1/3 h-2/3 bg-slate-800 p-5">
        <form method="POST" on:submit={aggiungi}>
            <input
                type="text"
                placeholder="Type here"
                name="title"
                class="input input-bordered w-full"
                bind:value={taskTitle}
            />
            <div class="flex justify-center items-center p-3">
                <button class="btn btn-warning w-full">Aggiungi</button>
            </div>
        </form>
    </div>

    <div class="w-1/3 h-2/3 bg-slate-800 p-5 overflow-scroll">
        {#if tasks.length == 0}
            <div>Non sono presenti task</div>
        {:else}
            {#each tasks as task}
                {#if task.status == "done"}
                    <div
                        class="glass w-full rounded-lg p-5 text-white font-bold text mb-2 line-through flex justify-between items-center"
                        data-task-id={task.id}
                    >
                        <div>
                            {task.title}
                        </div>
                        <div class="flex gap-1">
                            <button
                                class="btn bg-red-600 text-white"
                                on:click={rimuovi}
                                ><svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke-width="1.5"
                                    stroke="currentColor"
                                    class="size-6"
                                >
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"
                                    />
                                </svg></button
                            >
                        </div>
                    </div>
                {:else}
                    <div
                        class=" bg-slate-600 outline outline-white w-full rounded-lg p-5 text-white font-bold mb-2 flex justify-between items-center"
                        data-task-id={task.id}
                    >
                        <div>{task.title}</div>
                        <div class="flex gap-1">
                            <button
                                class="btn bg-red-600 text-white"
                                on:click={rimuovi}
                                ><svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke-width="1.5"
                                    stroke="currentColor"
                                    class="size-6"
                                >
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0"
                                    />
                                </svg>
                            </button>
                            <button
                                class="btn btn-success text-white"
                                on:click={aggiorna}
                                ><svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    fill="none"
                                    viewBox="0 0 24 24"
                                    stroke-width="1.5"
                                    stroke="currentColor"
                                    class="size-6"
                                >
                                    <path
                                        stroke-linecap="round"
                                        stroke-linejoin="round"
                                        d="m4.5 12.75 6 6 9-13.5"
                                    />
                                </svg>
                            </button>
                        </div>
                    </div>
                {/if}
            {/each}
        {/if}
    </div>
</div>
