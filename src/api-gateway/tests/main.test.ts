describe("eduflow", () => {
    const dir = "http://localhost:8080"
    test("main page", () => {
        fetch(dir + "/").then(response => {response.json().then(json => {
            expect(json).toEqual({
                "ok": true,
                "message": "",
                "data": "Hello, world!"
            });
        })})
    })

    test("info about me", () => {
        fetch(dir + "/my-user/", {
            headers: {
                "Authorization": "Bearer anqrfsNqOu"
            }
        }).then(response => {
            response.json().then(json => {
                expect(json).toEqual({
                    "ok": true,
                    "message": "",
                    "data": {
                        "id": 3,
                        "is_staff": false,
                        "username": "Coolman"
                    }
                })
        })})
    })
    test("sign in", () => {
        let data = new FormData()
        data.append("username", "abeme")
        data.append("password", "1234")
        fetch(dir + "/auth/sign-in/", {
            method: "POST",
            body: data,
        }).then(response => response.json().then(json => {
            expect(json).toEqual({
                "ok": true,
                "message": "",
                "data": {
                    "token": "abeme"
                }
            })
        }))
    })
})