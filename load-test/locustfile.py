from locust import HttpUser, task, between


class LoadUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def load(self):
        self.client.get("/load")
