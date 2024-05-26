import { Box, Button, Card, Flex, Heading, Input } from "@chakra-ui/react";
import { FormEvent, memo, useState } from "react";
import { useLoginMutation } from "../../redux/rtk/authApi";

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [loginMutation, { isLoading }] = useLoginMutation();

    const onSubmit = (e: FormEvent) => {
        e.preventDefault(); // Prevent default form submission behavior
        if (email === "" || password === "") {
            return;
        }
        loginMutation({ email, password });
    };

    return (
        <Flex
            align="center"
            justify="center"
            height="100vh"
            bg="gray.50"
            padding="4"
        >
            <Card
                width="full"
                maxW="sm"
                p="6"
                boxShadow="md"
                borderRadius="md"
                bg="white"
            >
                <form onSubmit={onSubmit}> {/* Form element */}
                    <Box mb="4">
                        <Heading size="md" textAlign="center" color="teal.500">
                            Login
                        </Heading>
                    </Box>
                    <Input
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        mb="4"
                        type="email"
                        variant="filled"
                        focusBorderColor="teal.400"
                    />
                    <Input
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        mb="4"
                        type="password"
                        variant="filled"
                        focusBorderColor="teal.400"
                    />
                    <Button
                        isLoading={isLoading}
                        colorScheme="teal"
                        width="full"
                        type="submit"
                        disabled={email === "" || password === ""}
                    >
                        Login
                    </Button>
                </form>
            </Card>
        </Flex>
    );
};

export default memo(Login);
