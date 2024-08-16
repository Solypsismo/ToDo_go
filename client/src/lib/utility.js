function updateStatusToDone(lis, id) {
    for (let i = 0; i < lis.length; i++) {
        if (lis[i].id === id) {
            lis[i].status = "done";
            return lis;
        }
    }
    return null;
}

function removeObj(lis, id) {
    for (let i = 0; i < lis.length; i++) {
        if (lis[i].id === id) {
            lis.splice(i, 1)
            return lis;
        }
    }
    return null;
}

export { updateStatusToDone, removeObj }