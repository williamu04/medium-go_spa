import { Box, Flex, Tabs } from "@chakra-ui/react";

const Stories = () => {
  return (
    <Flex>
      <Box>
        Articles
      </Box>
      <Tabs.Root defaultValue="draft" >
        <Tabs.List>
          <Tabs.Trigger value="draft">Draft</Tabs.Trigger>
          <Tabs.Trigger value="pubilsh">Published</Tabs.Trigger>
        </Tabs.List>
        <Tabs.Content value="draft" overflowY={"auto"}>Article yang masih draft</Tabs.Content>
        <Tabs.Content value="pubilsh">Article yang sudah dipublish</Tabs.Content>
      </Tabs.Root>
    </Flex>
  )
}

export default Stories;
