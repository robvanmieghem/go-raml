
def build_query_string(query_params=None):
    if query_params is None:
        return ""

    qs = "?"
    for key, elem in query_params:
        qs += key + "=" + str(elem) + "&"

    return qs[:-1]